package users

func (ul *UserList) get(id uint32) User {
	var res User
	for _, user := range *ul {
		if user.ID == id {
			res = user
			break
		}
	}
	return res
}

func (ul *UserList) add(u User) {

}

func (ul *UserList) update(id uint32, email string) User {
	var res User
	for i, user := range *ul {
		if user.ID == id {
			(*ul)[i].Email = email
			res = (*ul)[i]
			break
		}
	}

	return res
}

func (ul *UserList) delete(id uint32) {
	for i, user := range *ul {
		if user.ID == id {
			*ul = append((*ul)[:i], (*ul)[i+1:]...)
			break
		}
	}

}
