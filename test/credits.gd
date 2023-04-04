
extends VBoxContainer

func _ready():
	for t in get_children():
		if t is RichTextLabel:
			t.connect("meta_clicked", _on_meta_clicked)

func _on_meta_clicked(meta : String):
	OS.shell_open(meta)

