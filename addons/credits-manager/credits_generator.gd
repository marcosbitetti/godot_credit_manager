@tool
extends Window

const Helpers = preload("res://addons/credits-manager/helpers.gd")

const script_template = """
extends VBoxContainer

func _ready():
	for t in get_children():
		if t is RichTextLabel:
			t.connect("meta_clicked", _on_meta_clicked)

func _on_meta_clicked(meta : String):
	OS.shell_open(meta)

"""

func _on_generate_pressed():
	#save("res://test_folder")
	var d = EditorFileDialog.new()
	d.file_mode = EditorFileDialog.FILE_MODE_OPEN_DIR
	d.connect("dir_selected", dir_selected)
	add_child(d) # get_tree().root.
	d.popup_centered(Vector2(500,400))

func dir_selected(path : String):
	save(path)

func save(folder : String):
	var scene : VBoxContainer = VBoxContainer.new()
	scene.size_flags_horizontal = Control.SIZE_FILL
	scene.size_flags_vertical = Control.SIZE_FILL
	scene.custom_minimum_size = Vector2(800, 200)
	scene.name = "credits"
	
	var file = FileAccess.open(folder + "/credits.gd",FileAccess.WRITE)
	file.store_string(script_template)
	file.close()
	scene.set_script(load(folder + "/credits.gd"))
	
	var session = ""
	var rich : RichTextLabel
	for credit in Helpers.map_resources():
		# add label
		if credit.type != session:
			session = credit.type
			var label = Label.new()
			label.text = session
			label.name = "lbl_" + session
			scene.add_child(label)
			label.owner = scene
			rich = null
		if rich == null:
			rich = RichTextLabel.new()
			rich.size_flags_horizontal = Control.SIZE_EXPAND_FILL
			rich.size_flags_vertical = Control.SIZE_EXPAND_FILL
			rich.fit_content = true
			rich.scroll_active = false
			rich.shortcut_keys_enabled = false
			rich.bbcode_enabled = true
			rich.name = session
			scene.add_child(rich)
			rich.owner = scene
		var bbcode = Helpers.format_bb_code(credit)
		rich.text += bbcode + "\n\n"
	var packer = PackedScene.new()
	packer.pack(scene)
	
	if OK != ResourceSaver.save(packer, folder + "/credits.tscn", ResourceSaver.FLAG_NONE):
		OS.alert("cant save scene", "Error!")
	queue_free()


func _on_close_requested():
	queue_free()
