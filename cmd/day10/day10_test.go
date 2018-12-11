package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	points := []*Point{
		{
			position: &Vector{9, 1},
			velocity: &Vector{0, 2},
		},
		{
			position: &Vector{7, 0},
			velocity: &Vector{-1, 0},
		},
		{
			position: &Vector{3, -2},
			velocity: &Vector{-1, 1},
		},
		{
			position: &Vector{6, 10},
			velocity: &Vector{-2, -1},
		},
		{
			position: &Vector{2, -4},
			velocity: &Vector{2, 2},
		},
		{
			position: &Vector{-6, 10},
			velocity: &Vector{2, -2},
		},
		{
			position: &Vector{1, 8},
			velocity: &Vector{1, -1},
		},
		{
			position: &Vector{1, 7},
			velocity: &Vector{1, 0},
		},
		{
			position: &Vector{-3, 11},
			velocity: &Vector{1, -2},
		},
		{
			position: &Vector{7, 6},
			velocity: &Vector{-1, -1},
		},
		{
			position: &Vector{-2, 3},
			velocity: &Vector{1, 0},
		},
		{
			position: &Vector{-4, 3},
			velocity: &Vector{2, 0},
		},
		{
			position: &Vector{10, -3},
			velocity: &Vector{-1, 1},
		},
		{
			position: &Vector{5, 11},
			velocity: &Vector{1, -2},
		},
		{
			position: &Vector{4, 7},
			velocity: &Vector{0, -1},
		},
		{
			position: &Vector{8, -2},
			velocity: &Vector{0, 1},
		},
		{
			position: &Vector{15, 0},
			velocity: &Vector{-2, 0},
		},
		{
			position: &Vector{1, 6},
			velocity: &Vector{1, 0},
		},
		{
			position: &Vector{8, 9},
			velocity: &Vector{0, -1},
		},
		{
			position: &Vector{3, 3},
			velocity: &Vector{-1, 1},
		},
		{
			position: &Vector{0, 5},
			velocity: &Vector{0, -1},
		},
		{
			position: &Vector{-2, 2},
			velocity: &Vector{2, 0},
		},
		{
			position: &Vector{5, -2},
			velocity: &Vector{1, 2},
		},
		{
			position: &Vector{1, 4},
			velocity: &Vector{2, 1},
		},
		{
			position: &Vector{-2, 7},
			velocity: &Vector{2, -2},
		},
		{
			position: &Vector{3, 6},
			velocity: &Vector{-1, -1},
		},
		{
			position: &Vector{5, 0},
			velocity: &Vector{1, 0},
		},
		{
			position: &Vector{-6, 0},
			velocity: &Vector{2, 0},
		},
		{
			position: &Vector{5, 9},
			velocity: &Vector{1, -2},
		},
		{
			position: &Vector{14, 7},
			velocity: &Vector{-2, 0},
		},
		{
			position: &Vector{-3, 6},
			velocity: &Vector{2, -1},
		},
	}
	g := &Grid{points}

	fmt.Println(g.Draw())
}
