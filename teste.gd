extends BoxContainer

const Helpers = preload("res://addons/credits-manager/helpers.gd")

func _ready():
	for session in Helpers.map_resources():
		pass
