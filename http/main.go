package main

// https://api.eazy-mock.teqn.asia/mock/231084de-7e22-4f84-ba82-bb5e0f95431d/195/users/login

func main() {
	// r := http.Request{}
	// type name struct {
	// 	resp http.ResponseWriter
	// }
	// n := name{}

	// str := "{'name': ''bang'}"
	// n.resp.WriteHeader(http.StatusBadRequest)
	// n.resp.Write([]byte(str)) // write to response API
	// n.resp.Header().Add("mydata", "data")
	// type LoginRequest struct {
	// 	AccessToken string `json:"access_token"`
	// 	Provider    string `json:"provider"`
	// }

	// lr := LoginRequest{
	// 	AccessToken: "fake",
	// 	Provider:    "google",
	// }
	// body, err := json.Marshal(lr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defaultHost := `https://api.eazy-mock.teqn.asia/mock/231084de-7e22-4f84-ba82-bb5e0f95431d/195/users/login/google`
	// path := fmt.Sprintf("%s?id=%d&id2=%d", defaultHost, 1, 2)

	// ?id=2&id2=3

	// provider = 2
	// v1/user/2/info
	// tra ve 10 dong
	// music
	// swim
	// ...

	// set query params

	// v := url.Values{}
	// v.Add("id", "1")
	// v.Add("id2", "2")

	// str := v.Encode()
	// log.Println(str)

	// send a file
	// file, err := os.Open("ss")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// req, err := http.NewRequest(http.MethodPost,
	// 	path, file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Add("Content-Type", "binary/octet-stream")
	// req.Header.Add("Host", "test")
	// req.Header.Add("User-Agent", "test")

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	msg, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatalf("failed to read body, error: %v", err)
	// 	}
	// 	log.Fatalf("login failed with status: %d and message: %s", resp.StatusCode, string(msg))
	// }

	// var loginResponse map[string]interface{}
	// if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
	// 	log.Fatalf("failed to decode data, error: %v", err)
	// }

	// fmt.Println(loginResponse)
}
