# To-do

- [ ] Remake everything with new knowledge

- [ ] When making user CRUD, try do apopt a double hash approach:
1. First hash with SHA512 or something like that
2. Hash the SHA512 Hash with Bcrypt:
``` go
func main() {
	passwd := "pa55word"
	shaHash := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(passwd))
	fmt.Println("Token: ", shaHash)
	hash := sha256.Sum256([]byte(shaHash))
	final, err := bcrypt.GenerateFromPassword(hash[:], bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(final))
	fmt.Println(bcrypt.CompareHashAndPassword(final, hash[:]))
}
```

- [ ] Relationship querys using innerjoins
