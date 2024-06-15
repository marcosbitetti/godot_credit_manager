extends Object


static func get_from_api(args : Array[String]) -> Array:
	var path_base : String = ProjectSettings.globalize_path("res://")
	var path = path_base + "addons/credits-manager/aux/credits_handler/bin/" + check_os()
	var path_handler = path_base + "addons/credits-manager/aux/credits_handler"
	var out : Array[String] = []
	OS.set_environment('ATRIBUITION_HANDLER_PATH', ProjectSettings.globalize_path(path_handler))
	var code : int = OS.execute(path, args, out, true )
	if code != 0:
		print_debug("error: " + str(code))
		print_debug(out)
		return []
	if out.size() == 0:
		print_debug("error on call binary")
		return []
	var list = JSON.parse_string(out[0])
	if not list is Array:
		list = [list]
	return Array(list)

static func format_bb_code(credit : Dictionary) -> String:
	return '[b][url=' + credit.link + ']' + credit.name + '[/url][/b]' +\
		' by [b]' + credit.author + '[/b],  licensed under ' +\
		'[b][url=' + credit.licenceUrl + ']' + credit.licence + '[/url][/b]'

static func map_resources() -> Array:
	var local = files_resources()
	var list_base = []
	for c in get_from_api(["list", "asc"]):
		if local.find(c.filename) > -1:
			list_base.append(c)
	
	var list_final = []
	for session_obj in get_from_api(["types","asc"]):
		for item in list_base:
			if item.type == session_obj.name:
				list_final.append(item)
	return list_final

static func files_resources(path : String = "", list : Array = []) -> Array:
	path = "res://" if "" == path else path
	var dir = DirAccess.open(path)
	if not dir:
		OS.alert("An error occurred when trying to access the path", "Error!")
		return []
	else:
		dir.list_dir_begin()
		var file_name = dir.get_next()
		while file_name != "":
			if dir.current_is_dir():
				if file_name != "." and file_name != "..":
					list.append(file_name)
					files_resources(path + "/" + file_name, list)
			else:
				list.append(file_name)
			file_name = dir.get_next()
	return list
	
static func check_os() -> String:
	var run = ""
	match OS.get_name():
		"Windows", "UWP":
			run = "credits-manager-amd64.exe"
		"macOS":
			run = "credits-manager-amd64-darwin"
		"Linux", "FreeBSD", "NetBSD", "OpenBSD", "BSD":
			run = "credits-manager-amd64-linux"
	return run


