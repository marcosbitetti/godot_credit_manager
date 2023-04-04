@tool
extends VBoxContainer

const Helpers = preload("res://addons/credits-manager/helpers.gd")
const CreditItem : PackedScene = preload("res://addons/credits-manager/credit_item.tscn")
const CreditEdit : PackedScene = preload("res://addons/credits-manager/credit_edit.tscn")
const CreditGenerator : PackedScene = preload("res://addons/credits-manager/credits_generator.tscn")

func update_credits():
	while $scroll/list.get_child_count():
		var n = $scroll/list.get_child(0)
		$scroll/list.remove_child(n)
		n.queue_free()
	for c in Helpers.get_from_api(["list", "asc"]):
		var item = CreditItem.instantiate()
		item.parse(c)
		item.container(self)
		$scroll/list.add_child(item)
		$scroll/list.add_child( HSeparator.new() )

# Called when the node enters the scene tree for the first time.
func _ready():
	update_credits()


func _on_add_pressed():
	var c = CreditEdit.instantiate()
	get_tree().root.add_child(c)
	c.popup_centered(c.size)
	c.setup({"_id":0,"author":"","filename":"","licence":"","licenceUrl":"","link":"","name":"","type":""})
	c.connect("update_data", create_data)

func create_data():
	update_credits()


func _on_export_json_pressed():
	pass # Replace with function body.


func _on_export_tscene_pressed():
	var c = CreditGenerator.instantiate()
	get_tree().root.add_child(c)
	c.popup_centered(c.size)
