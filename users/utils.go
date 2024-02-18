package users

func (ul *UserList) get(id int64) User {
	var res User
	for _, user := range *ul {
		if user.Id == id {
			res = user
			break
		}
	}
	return res
}

func (ul *UserList) add(u User) {
	*ul = append(*ul, u)
}

func (ul *UserList) update(id int64, email string) User {
	var res User
	for i, user := range *ul {
		if user.Id == id {
			(*ul)[i].Email = email
			res = (*ul)[i]
			break
		}
	}

	return res
}

func (ul *UserList) delete(id int64) {
	for i, user := range *ul {
		if user.Id == id {
			*ul = append((*ul)[:i], (*ul)[i+1:]...)
			break
		}
	}

}
