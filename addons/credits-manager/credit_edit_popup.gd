@tool
extends Window

signal update_data()

const Helpers = preload("res://addons/credits-manager/helpers.gd")

var data : Dictionary
var types : Array[Dictionary] = []
var licences : Array[Dictionary] = []

func setup(_data : Dictionary):
	data = _data
	$Panel/a/name/LineEdit.text = data.name
	$Panel/a/link/LineEdit.text = data.link
	$Panel/a/file/LineEdit.text = data.filename
	$Panel/a/author/LineEdit.text = data.author
	$Panel/a/type/MenuButton.text = data.type
	$Panel/a/licence/MenuButton.text = data.licence
	

func _on_close_requested():
	queue_free()


func _on_type_about_to_popup():
	var pop : PopupMenu = $Panel/a/type/MenuButton.get_popup()
	if types.size() == 0 :
		for i in Helpers.get_from_api(["types", "asc"]):
			types.append(i)
			pop.add_item(i.name)
		pop.connect("index_pressed", _on_change_type)
	pop.popup()


func _on_licence_about_to_popup():
	var pop : PopupMenu = $Panel/a/licence/MenuButton.get_popup()
	if licences.size() == 0 :
		for i in Helpers.get_from_api(["licences", "asc"]):
			licences.append(i)
			pop.add_item(i.name)
		pop.connect("index_pressed", _on_change_licence)
	pop.popup()


func _on_change_type(index : int):
	data.type = types[index].name
	$Panel/a/type/MenuButton.text = data.type
	
func _on_change_licence(index : int):
	data.licence = licences[index].name
	data.licenceUrl = licences[index].link
	$Panel/a/licence/MenuButton.text = data.licence



func _on_button_save_pressed():
	Helpers.get_from_api(["update", JSON.stringify(data).replace("\"", "\\\"")])
	emit_signal("update_data")
	call_deferred("queue_free")


func _on_name_text_changed(new_text):
	data.name = new_text
	

func _on_link_text_changed(new_text):
	data.link = new_text


func _on_file_text_changed(new_text):
	data.filename = new_text


func _on_author_text_changed(new_text):
	data.author = new_text
	

func _on_pick_pressed():
	var fd : EditorFileDialog = EditorFileDialog.new()
	fd.file_mode = EditorFileDialog.FILE_MODE_OPEN_ANY
	fd.connect("file_selected", file_selected)
	fd.connect("dir_selected", file_selected)
	get_tree().root.add_child(fd)
	fd.popup_centered()

func file_selected(path : String):
	var name = path.get_file()
	if name=="":
		name = path.get_slice("/", path.get_slice_count("/") - 1)
	data.filename = name
	$Panel/a/file/LineEdit.text = data.filename

