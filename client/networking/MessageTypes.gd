###############################################################
# This code is generated by `go run ./scripts/generate.go`.   #
# If you want to generate the contents of this file, go to    #
# the server directory and run the command.                   #
#                                                             #
# DON'T TOUCH IT DIRECTLY! YOU WILL SUFFER!                   #
###############################################################

extends Node

const KEEP_ALIVE = 100
const ACTION_ERROR = 101
const AUTHENTICATE = 200
const CREATE_LOBBY = 201
const CREATED_LOBBY_RESPONSE = 301
const DELETE_LOBBY = 202
const DELETE_LOBBY_RESPONSE = 302
const LIST_LOBBIES = 203
const LIST_LOBBIES_RESPONSE = 303
const JOIN_LOBBY = 204
const JOIN_LOBBY_RESPONSE = 304
const PLAYER_LOBBY_JOINED = 305
const LIST_LOBBY_PLAYERS = 206
const LIST_LOBBY_PLAYERS_RESPONSE = 306
const START_GAME = 207
const START_GAME_RESPONSE = 307
const LOBBY_PLAYER_CONNECTED = 308
const LOBBY_PLAYER_DISCONNECTED = 309
const GAME_END = 310
const GAME_RECONNECT_AVAILABLE = 211
const GAME_RECONNECT_AVAILABLE_RESPONSE = 311
const RECONNECT = 212
const RECONNECT_RESPONSE = 312
const LEAVE_GAME = 213
const LEAVE_GAME_RESPONSE = 313
const LEAVE_LOBBY = 214
const LEAVE_LOBBY_RESPONSE = 314
const PLAYER_MOVE = 400
const SHOOT_PROJECTILE = 401
const UPDATE_STATE = 500
const PLAYER_DISCONNECTED = 501
const PLAYER_CONNECTED = 502
