package main

import "gopkg.in/src-d/go-git.v4"

const daysInLastSixMonths = 183

func ProcessRepositories(email string) map[int]int {

	filepath := getDotFilePath()
	repos := parseFileLinesToSlice(filepath)
	daysInMap := daysInLastSixMonths

	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0
	}

	for _, path := range repos {
		commits = fileCommits(email, path, commits)
	}

	return commits
}

func fileCommits(email string, path string, commits map[int]int) map[int]int {
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	ref, err := repo.Head()

	if err != nil {
		panic(err)
	}

	iterator, errr := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}
	offset := calcOffset()
	err = iterator.ForEach()
}
