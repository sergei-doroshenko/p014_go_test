package controllers

import "github.com/sergei-doroshenko/p014_go_test/models"

var posts map[string]*models.Post = make(map[string]*models.Post, 0)
var postsMD map[string]*models.PostMD = make(map[string]*models.PostMD, 0)

var counter int32
