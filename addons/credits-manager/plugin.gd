@tool
extends EditorPlugin

var credit
var _menu = MenuButton.new()
var help
var about


func _enter_tree():
	credit = preload("res://addons/credits-manager/credits.tscn").instantiate()
	add_control_to_bottom_panel(credit, "Credits Manager")
	
func _exit_tree():
	remove_control_from_bottom_panel(credit)
	var old = credit
	credit = null
	old.queue_free()


func item_selected(id):
	pass
	

func handles(obj):
	if _menu.is_inside_tree():
		return true
		
	# add_control_to_container( CONTAINER_CANVAS_EDITOR_MENU, _menu )
	
	return true
