package Proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %s could not be found\n", id)
}

// Implementasi dari UserFinder
type UserListProxy struct {
	// seolah-olah database untuk mencari user
	SomeDatabase UserList

	// merupakan stack yang menyimpan user yang diakses
	StackCache UserList

	// Jumlah maksimal user yang disimpan dalam stack
	StackCapacity int

	// variabel untuk pemberitahuan kalau hasil searching user terakhir disimpan cache
	DidDidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {

	/*
		Pencarian menggunakan proxy.

		yang mencari terlihat adalah
		stack cache, benar memang tapi
		kalau dilihat lebih dalam yang
		mencari adalah UserList yang
		sebenarnya dibungkus oleh Proxy
	*/
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	fmt.Println("Returning user from database")
	u.DidLastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}
