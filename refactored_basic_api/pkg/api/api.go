package api

/*
var Articles []m.Article


func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article m.Article
	json.Unmarshal(reqBody, &article)
	article.Id = uuid.New()
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
	vars := mux.Vars(r)

	key, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Invalid key format, %s", vars["id"])
		w.WriteHeader(500)
		fmt.Fprintf(w, "Invalid key format %s", vars["id"])
		return
	}
	for index, article := range Articles {
		log.Printf("article.Id %s, key %s", article.Id, key)
		if article.Id == key {
			log.Printf("found a key %s", key)
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	vars := mux.Vars(r)
	key, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Invalid key format, %s", vars["id"])
		w.WriteHeader(500)
		fmt.Fprintf(w, "Invalid key format %s", vars["id"])
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article m.Article
	json.Unmarshal(reqBody, &article)

	for index, a := range Articles {
		log.Printf("article.Id %s, key %s", article.Id, key)
		if a.Id == key {
			log.Printf("found a key %s", key)
			a.Title = article.Title
			a.Desc = article.Desc
			a.Content = article.Content
			Articles[index] = a
		}
	}
}

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}*/
