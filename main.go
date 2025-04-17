package main

import (
	"encoding/json"
	"net/http"
)

type Address struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Accuracy  string  `json:"accuracy"`
	} `json:"coordinates"`
	Neighborhood struct {
		Name       string `json:"name"`
		Population int    `json:"population"`
		Landmarks  []struct {
			Name  string            `json:"name"`
			Type  string            `json:"type"`
			Hours map[string]string `json:"hours"`
		} `json:"landmarks"`
	} `json:"neighborhood"`
}

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Address     Address `json:"address"`
	Preferences struct {
		Language      string `json:"language"`
		Timezone      string `json:"timezone"`
		Notifications struct {
			Email       bool `json:"email"`
			SMS         bool `json:"sms"`
			Push        bool `json:"push"`
			Preferences struct {
				Email struct {
					Frequency string   `json:"frequency"`
					Topics    []string `json:"topics"`
				} `json:"email"`
				SMS struct {
					Frequency string   `json:"frequency"`
					Topics    []string `json:"topics"`
				} `json:"sms"`
			} `json:"preferences"`
		} `json:"notifications"`
	} `json:"preferences"`
	Subscriptions []struct {
		PlanName  string   `json:"plan_name"`
		Validity  string   `json:"validity"`
		StartDate string   `json:"start_date"`
		EndDate   string   `json:"end_date"`
		AutoRenew bool     `json:"auto_renew"`
		Features  []string `json:"features"`
	} `json:"subscriptions"`
	Friends []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Contact struct {
			Email   string `json:"email"`
			Phone   string `json:"phone"`
			Address struct {
				Street  string `json:"street"`
				City    string `json:"city"`
				State   string `json:"state"`
				ZipCode string `json:"zip_code"`
			} `json:"address"`
		} `json:"contact"`
		Preferences struct {
			Language      string `json:"language"`
			Timezone      string `json:"timezone"`
			Notifications struct {
				Email bool `json:"email"`
				SMS   bool `json:"sms"`
			} `json:"notifications"`
		} `json:"preferences"`
	} `json:"friends"`
	Employment struct {
		Status  string `json:"status"`
		Company struct {
			Name     string `json:"name"`
			Industry string `json:"industry"`
			Location struct {
				City    string `json:"city"`
				State   string `json:"state"`
				ZipCode string `json:"zip_code"`
			} `json:"location"`
			Positions []struct {
				Title      string `json:"title"`
				Department string `json:"department"`
				StartDate  string `json:"start_date"`
				EndDate    string `json:"end_date"`
				Salary     struct {
					Currency string  `json:"currency"`
					Amount   float64 `json:"amount"`
				} `json:"salary"`
			} `json:"positions"`
		} `json:"company"`
		FreelanceProjects []struct {
			ProjectName string `json:"project_name"`
			Client      string `json:"client"`
			StartDate   string `json:"start_date"`
			EndDate     string `json:"end_date"`
			Payment     struct {
				Currency string  `json:"currency"`
				Amount   float64 `json:"amount"`
			} `json:"payment"`
		} `json:"freelance_projects"`
	} `json:"employment"`
	AuditLogs []struct {
		Timestamp string `json:"timestamp"`
		Action    string `json:"action"`
		Details   struct {
			IPAddress string `json:"ip_address"`
			Device    string `json:"device"`
			Location  string `json:"location"`
		} `json:"details"`
	} `json:"audit_logs"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Create a sample nested JSON response
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			ZipCode: "12345",
		},
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the user struct to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func userHandler2(w http.ResponseWriter, r *http.Request) {
	// Create a sample nested JSON response
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			ZipCode: "12345",
			Coordinates: struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Accuracy  string  `json:"accuracy"`
			}{
				Latitude: 37.7749, Longitude: -122.4194, Accuracy: "high",
			},
			Neighborhood: struct {
				Name       string `json:"name"`
				Population int    `json:"population"`
				Landmarks  []struct {
					Name  string            `json:"name"`
					Type  string            `json:"type"`
					Hours map[string]string `json:"hours"`
				} `json:"landmarks"`
			}{
				Name: "Downtown", Population: 5000,
				Landmarks: []struct {
					Name  string            `json:"name"`
					Type  string            `json:"type"`
					Hours map[string]string `json:"hours"`
				}{
					{
						Name: "City Hall",
						Type: "Government Building",
						Hours: map[string]string{
							"monday":    "9 AM - 5 PM",
							"tuesday":   "9 AM - 5 PM",
							"wednesday": "9 AM - 5 PM",
						},
					},
					{
						Name: "Central Park",
						Type: "Public Park",
						Hours: map[string]string{
							"monday":    "6 AM - 10 PM",
							"tuesday":   "6 AM - 10 PM",
							"wednesday": "6 AM - 10 PM",
						},
					},
				},
			},
		},
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the user struct to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func testingHandler2(w http.ResponseWriter, r *http.Request) {
	// Directly define the JSON string to be returned
	jsonResponse := `{
		"id": 1,
		"name": "John Doe",
		"email": "john.doe@example.com",
		"address": {
			"street": "123 Main St",
			"city": "Anytown",
			"state": "CA",
			"zip_code": "12345",
			"coordinates": {
				"latitude": 37.7749,
				"longitude": -122.4194,
				"accuracy": "high"
			},
			"neighborhood": {
				"name": "Downtown",
				"population": 5000,
				"landmarks": [
					{
						"name": "City Hall",
						"type": "Government Building",
						"hours": {
							"monday": "9 AM - 5 PM",
							"tuesday": "9 AM - 5 PM",
							"wednesday": "9 AM - 5 PM",
							"thursday": "9 AM - 5 PM",
							"friday": "9 AM - 5 PM"
						}
					},
					{
						"name": "Central Park",
						"type": "Public Park",
						"hours": {
							"monday": "6 AM - 10 PM",
							"tuesday": "6 AM - 10 PM",
							"wednesday": "6 AM - 10 PM",
							"thursday": "6 AM - 10 PM",
							"friday": "6 AM - 10 PM",
							"saturday": "8 AM - 11 PM",
							"sunday": "8 AM - 11 PM"
						}
					}
				]
			}
		}
	}`

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON string directly to the response
	w.Write([]byte(jsonResponse))
}

func testingHandler(w http.ResponseWriter, r *http.Request) {
	// Directly define the JSON string to be returned
	jsonResponse := `[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "address": {
            "street": "123 Main St",
            "city": "Anytown",
            "state": "CA",
            "zip_code": "12345",
            "coordinates": {
                "latitude": 37.7749,
                "longitude": -122.4194,
                "accuracy": "high"
            },
            "neighborhood": {
                "name": "Downtown",
                "population": 5000,
                "landmarks": [
                    {
                        "name": "City Hall",
                        "type": "Government Building",
                        "hours": {
                            "monday": "9 AM - 5 PM",
                            "tuesday": "9 AM - 5 PM",
                            "wednesday": "9 AM - 5 PM"
                        }
                    },
                    {
                        "name": "Central Park",
                        "type": "Public Park",
                        "hours": {
                            "monday": "6 AM - 10 PM",
                            "tuesday": "6 AM - 10 PM",
                            "wednesday": "6 AM - 10 PM"
                        }
                    }
                ]
            }
        },
        "preferences": {
            "language": "en",
            "timezone": "PST",
            "notifications": {
                "email": true,
                "sms": false,
                "push": true,
                "preferences": {
                    "email": {
                        "frequency": "daily",
                        "topics": ["updates", "offers", "news"]
                    },
                    "sms": {
                        "frequency": "weekly",
                        "topics": ["promotions"]
                    }
                }
            }
        },
        "subscriptions": [
            {
                "plan_name": "Premium",
                "validity": "1 Year",
                "start_date": "2023-01-01",
                "end_date": "2023-12-31",
                "auto_renew": true,
                "features": ["Unlimited Access", "Priority Support", "Advanced Analytics"]
            },
            {
                "plan_name": "Standard",
                "validity": "6 Months",
                "start_date": "2023-06-01",
                "end_date": "2023-11-30",
                "auto_renew": false,
                "features": ["Basic Access", "Standard Support"]
            }
        ],
        "friends": [
            {
                "id": 2,
                "name": "Jane Smith",
                "contact": {
                    "email": "jane.smith@example.com",
                    "phone": "+1234567890",
                    "address": {
                        "street": "456 Elm St",
                        "city": "Othertown",
                        "state": "CA",
                        "zip_code": "67890",
                        "coordinates": {
                            "latitude": 34.0522,
                            "longitude": -118.2437,
                            "accuracy": "medium"
                        },
                        "neighborhood": {
                            "name": "Uptown",
                            "population": 3000,
                            "landmarks": [
                                {
                                    "name": "Uptown Library",
                                    "type": "Public Library",
                                    "hours": {
                                        "monday": "10 AM - 6 PM",
                                        "tuesday": "10 AM - 6 PM",
                                        "wednesday": "10 AM - 6 PM"
                                    }
                                },
                                {
                                    "name": "Riverside Park",
                                    "type": "Public Park",
                                    "hours": {
                                        "monday": "5 AM - 11 PM",
                                        "tuesday": "5 AM - 11 PM",
                                        "wednesday": "5 AM - 11 PM"
                                    }
                                }
                            ]
                        }
                    }
                },
                "preferences": {
                    "language": "es",
                    "timezone": "CET",
                    "notifications": {
                        "email": false,
                        "sms": true
                    }
                }
            },
            {
                "id": 3,
                "name": "Mark Johnson",
                "contact": {
                    "email": "mark.johnson@example.com",
                    "phone": "+0987654321",
                    "address": {
                        "street": "789 Oak St",
                        "city": "Newcity",
                        "state": "NY",
                        "zip_code": "98765",
                        "coordinates": {
                            "latitude": 40.7128,
                            "longitude": -74.006,
                            "accuracy": "high"
                        },
                        "neighborhood": {
                            "name": "Midtown",
                            "population": 4000,
                            "landmarks": [
                                {
                                    "name": "Midtown Museum",
                                    "type": "Cultural Institution",
                                    "hours": {
                                        "monday": "9 AM - 5 PM",
                                        "tuesday": "9 AM - 5 PM",
                                        "wednesday": "9 AM - 5 PM"
                                    }
                                },
                                {
                                    "name": "Grand Central Park",
                                    "type": "Public Park",
                                    "hours": {
                                        "monday": "6 AM - 10 PM",
                                        "tuesday": "6 AM - 10 PM",
                                        "wednesday": "6 AM - 10 PM"
                                    }
                                }
                            ]
                        }
                    }
                },
                "preferences": {
                    "language": "en",
                    "timezone": "EST",
                    "notifications": {
                        "email": true,
                        "sms": true
                    }
                }
            }
        ],
        "employment": {
            "status": "Employed",
            "company": {
                "name": "Tech Innovations Inc.",
                "industry": "Software Development",
                "location": {
                    "city": "San Francisco",
                    "state": "CA",
                    "zip_code": "94105"
                },
                "positions": [
                    {
                        "title": "Software Engineer",
                        "department": "Engineering",
                        "start_date": "2020-03-01",
                        "end_date": null,
                        "salary": {
                            "currency": "USD",
                            "amount": 120000
                        }
                    }
                ]
            },
            "freelance_projects": [
                {
                    "project_name": "Web App Development",
                    "client": "XYZ Corp",
                    "start_date": "2023-01-15",
                    "end_date": "2023-06-15",
                    "payment": {
                        "currency": "USD",
                        "amount": 25000
                    }
                }
            ]
        },
        "audit_logs": [
            {
                "timestamp": "2024-01-01T12:00:00Z",
                "action": "Login",
                "details": {
                    "ip_address": "192.168.1.1",
                    "device": "Laptop",
                    "location": "San Francisco, CA"
                }
            },
            {
                "timestamp": "2024-01-02T14:30:00Z",
                "action": "Password Change",
                "details": {
                    "ip_address": "192.168.1.2",
                    "device": "Mobile Phone",
                    "location": "Los Angeles, CA"
                }
            }
        ]
    },
]
`

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON string directly to the response
	w.Write([]byte(jsonResponse))
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/testing", testingHandler)
	http.HandleFunc("/testing2", testingHandler2)
	http.HandleFunc("/user2", userHandler2)
	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

