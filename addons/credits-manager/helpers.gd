extends Object


static func get_from_api(args : Array[String]) -> Array:
	var path : String = ProjectSettings.globalize_path("res://")
	path += "addons/credits-manager/aux/credits_handler/bin/credits-manager-amd64-linux"
	var out : Array[String] = []
	var code : int = OS.execute(path, args, out, true )
	if code != 0:
		print_debug("error: " + str(code))
		print_debug(out)
		return []
	if out.size() == 0:
		print_debug("error on call binary")
		return []
	var list = JSON.parse_string(out[0])
	return Array(list)

static func format_bb_code(credit : Dictionary) -> String:
	return '[b][url=' + credit.link + ']' + credit.name + '[/url][/b]' +\
		' by [b]' + credit.author + '[/b],  licensed under ' +\
		'[b][url=' + credit.licenceUrl + ']' + credit.licence + '[/url][/b]'

static func map_resources() -> Array:
	var local = files_resources()
	var list = []
	for c in get_from_api(["list", "asc"]):
		if local.find(c.filename) > -1:
			list.append(c)
	return list

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
	
