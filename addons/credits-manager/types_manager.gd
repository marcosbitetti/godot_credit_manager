@tool
extends Window

const Helpers = preload("res://addons/credits-manager/helpers.gd")

var id_edit : int = 0

func update_list():
	while $VBoxContainer/ScrollContainer/list.get_child_count():
		var n = $VBoxContainer/ScrollContainer/list.get_child(0)
		$VBoxContainer/ScrollContainer/list.remove_child(n)
		n.queue_free()
	var command = "asc"
	for l in Helpers.get_from_api(["types", command]):
		var bt = Button.new()
		bt.size_flags_horizontal = Control.SIZE_EXPAND_FILL
		bt.text = l.name
		bt.alignment = HORIZONTAL_ALIGNMENT_LEFT
		bt.set_meta("_id", l._id)
		$VBoxContainer/ScrollContainer/list.add_child(bt)
		bt.connect("pressed", func(): edit(bt))

func _ready():
	update_list()

func edit(bt : Button):
	id_edit = bt.get_meta("_id")
	$VBoxContainer/edit/HBoxContainer/name.text = bt.text
	$VBoxContainer/edit/HBoxContainer3/HBoxContainer2/Button.text = tr("Update")


func _on_button_2_pressed():
	id_edit = 0
	$VBoxContainer/edit/HBoxContainer/name.text = ""
	$VBoxContainer/edit/HBoxContainer3/HBoxContainer2/Button.text = tr("Create New")


func _on_button_pressed():
	var data  = {}
	if $VBoxContainer/edit/HBoxContainer/name.text == "":
		return
	data.name = $VBoxContainer/edit/HBoxContainer/name.text
	if id_edit == 0:
		Helpers.get_from_api(["add-type", JSON.stringify(data).replace("\"", "\\\"")])
	else:
		data._id = id_edit
		Helpers.get_from_api(["update-type", JSON.stringify(data).replace("\"", "\\\"")])
	_on_button_2_pressed()
	update_list()


func _on_close_requested():
	queue_free()


func _on_button_del_pressed():
	var com : ConfirmationDialog = ConfirmationDialog.new()
	get_tree().root.add_child(com)
	com.title = "Confirm Exclusion"
	com.get_label().text = "Remove " + $VBoxContainer/edit/HBoxContainer/name.text + "?\n" +\
		"That actions has no undo."
	com.get_ok_button().connect("pressed", perform_delete)
	com.popup_centered()

func perform_delete():
	Helpers.get_from_api(["delete-type", str(id_edit)])
	_on_button_2_pressed()
	update_list()

