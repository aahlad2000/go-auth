<h1>Authentication application in GO</h1>

Simple application in GO to register users and create JWT token based on the credentials

<h2>Endpoints</h2>

/signup - To register

/login - To generate the JWT token

<h2>Libraries used:<h2>

gin - to support HTTP requests
bcrypt - to store the password hash
jwt - to create JWT tokens
gorm - for automatic table creation and db connection

<h2>Work In Progress</h2>
- To check and validate user
- Custom expiry date of JWT token
- To get user details from JWT token
- To cache the token and send the cached token for duplicate requests
