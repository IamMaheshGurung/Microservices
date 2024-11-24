
/dairy-shop
├── /cmd
│   ├── /user-service
│   │   └── main.go
│   ├── /product-service
│   │   └── main.go
│   └── /order-service
│       └── main.go
├── /services
│   ├── /user-service
│   │   ├── user_service.go
│   │   ├── user_handler.go
│   │   └── user_model.go
│   ├── /product-service
│   │   ├── product_service.go
│   │   ├── product_handler.go
│   │   └── product_model.go
│   ├── /order-service
│   │   ├── order_service.go
│   │   ├── order_handler.go
│   │   └── order_model.go
├── /frontend
│   ├── /static
│   │   ├── style.css
│   │   └── script.js
│   ├── /templates
│   │   ├── register.html
│   │   ├── otp_verification.html
│   │   └── product_list.html
│   └── /handlers
│       └── user_handlers.go
├── /config
│   ├── config.go
│   └── db_config.json
├── /test
│   ├── /product-service_test.go
│   ├── /order-service_test.go
│   └── /user-service_test.go
└── README.md

