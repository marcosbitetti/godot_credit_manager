@tool
extends HBoxContainer

const Helpers = preload("res://addons/credits-manager/helpers.gd")
const CreditEdit : PackedScene = preload("res://addons/credits-manager/credit_edit.tscn")

var data : Dictionary
var panel

func parse(_data : Dictionary):
	data = _data
	$text.parse_bbcode(Helpers.format_bb_code(data))

func container(nd):
	panel = nd

func _on_edit_pressed():
	var c = CreditEdit.instantiate()
	get_tree().root.add_child(c)
	c.popup_centered(c.size)
	c.setup(data)
	c.connect("update_data", update_data)


func _on_delete_pressed():
	var com : ConfirmationDialog = ConfirmationDialog.new()
	get_tree().root.add_child(com)
	com.title = "Confirm Exclusion"
	com.get_label().text = "Remove " + data.name + "?\n" +\
		"That actions has no undo."
	com.get_ok_button().connect("pressed", perform_delete)
	com.popup_centered()

func update_data():
	panel.update_credits()

func perform_delete():
	Helpers.get_from_api(["delete", str(data._id)])
	update_data()
