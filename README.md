# igslim

Get Instagram user profile.

Docs: <https://pkg.go.dev/github.com/caiguanhao/igslim>

```go
client := igslim.NewClient(os.Getenv("IGSESSIONID"))
user, err := client.GetUser("TaylorSwift")
if err != nil {
	panic(err)
}
enc := json.NewEncoder(os.Stdout)
enc.SetEscapeHTML(false)
enc.SetIndent("", "  ")
enc.Encode(user)
```

```
=== RUN   TestGetUser
{
  "Id": 11830955,
  "FbId": 17841401648650184,
  "UserName": "taylorswift",
  "FullName": "Taylor Swift",
  "Verified": true,
  "Picture": "https://scontent-hkt1-2.cdninstagram.com/v/t51.2885-19/s320x320/203390676_6325176860841825_830428569594643688_n.jpg?_nc_ht=scontent-hkt1-2.cdninstagram.com&_nc_ohc=il4wZ60N2tQAX9QjskR&tn=vkikvA8yTJv52AHZ&edm=ABfd0MgBAAAA&ccb=7-4&oh=14e473e45306615c7cf55ebb4afd27f5&oe=6150BAD1&_nc_sid=7bff83",
  "Biography": "Happy, free, confused and lonely at the same time. \nNov. 19th, 2021",
  "CategoryName": "Musician",
  "FollowingsCount": 0,
  "FollowersCount": 179511327,
  "PostsCount": 512
}
```
