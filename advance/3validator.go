type User struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"gte=18,lte=120"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Parse the JSON request and populate the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(user)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	// Validation passed, proceed to process the user data
	// Your application logic goes here...

	// Send a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User created successfully!")
}
