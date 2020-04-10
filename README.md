# Learn web application golang

-   https://itnext.io/beautify-your-golang-project-f795b4b453aa
-   https://adodd.net/post/go-ddd-repository-pattern/

-   Project structure

        ├── go.mod
        ├── main.go
        ├── project
        │   ├── errors.go
        │   ├── driver
        │   ├── handler
        │   ├── models
        │   └── repository
        └── README.md

-   driver - This layer is responsible for connecting with the database
    -   declare DB struct

            type DB struct {
                SQL *sql.DB
            }

    -   declare ConnectSQl func with return database connection

            func ConnectSQL(host string, port int64, username string , password string, dbname string) (*DB, error) {
                var dbConn = &DB{}

                dbSource := fmt.Sprintf(
                    "%s:%s@tcp(%s:%d)/%s?charset=utf8",
                    username,
                    password,
                    host,
                    port,
                    dbname,
                )

                dbConn.SQL, err := sql.Open("mysql", dbSource)

                if err != nil {
                    panic(err)
                }

                return dbConn, err
            }

-   models
    -   Store all models struct -> sql tables
    -   post.go

            type Post struct {
                ID int64 `json:"id"`
                Title string `json:"title"`
                Content string `json:"content"`
            }

-   repository
    -   Declare `PostRepo` interface to have needed methods
        -   Fetch       -> gets all Post from database, using HTTP GET method
        -   GetByID     -> get a Post by ID, using HTTP GET method
        -   Create      -> create a Post, using HTTP POST method
        -   Update      -> update a Post, using HTTP POST method
        -   Delete      -> delete a Post, using HTTP POST method
    -   `repository.go`

            type PostRepo interface {
                Fetch(pageNumber int64) ([]*models.Post, error)
                GetByID(id int) (*models.Post, error)
                Create(p *models.Post) (int64, error)
                Update(p *models.Post) (*models.Post, error)
                Delete(id int64) (bool, error)
            }

    -   Responsible for database job - no business logic is implemented
        -   Query
        -   Update
        -   Delete
        -   Insert
    -   `post/post.go`
        -   declare sqlPostRepo with has a db connection

                type sqlPostRepo struct {
                    Conn *sql.DB
                }

        -   declare function with return `repository.PostRepo`

                func NewSQLPostRepo(Conn *sql.DB) repository.PostRepo {
                    return &sqlPostRepo {
                        Conn: Conn,
                    }
                }

        -   implementing all PostRepo interface's methods

                func (m *sqlPostRepo) Fetch(pageNumber int64) ([]*models.Post, error)
                func (m *sqlPostRepo) GetByID(id int64) (*models.Post, error)
                func (m *sqlPostRepo) Create(p *models.Post) (int64, error)
                func (m *sqlPostRepo) Update(p *models.Post) (*models.Post, error)
                func (m *sqlPostRepo) Delete(id int64) (bool, error)

-   handler
    -   `handler/post.go` - Post controllers

            type Post struct {
                repo repository.PostRepo
            }

    -   init Post controller with database connection

            func NewPostHandler(db *driver.DB) *Post {
                return &Post{
                    repo: repository.NewSQLPostRepo(db.SQL)
                }
            }

    -   declare all functions of Post Controllers

            func (p *Post) Fetch(res http.ResponseWriter, req *http.Request)
            func (p *Post) GetByID(res http.ResponseWriter, req *http.Request)
            func (p *Post) Create(res http.ResponseWriter, req *http.Request)
            func (p *Post) Update(res http.ResponseWriter, req *http.Request)
            func (p *Post) Delete(res http.ResponseWriter, req *http.Request)

-   errors.go - error helpers
-   mysql
    -   user
        -   root/Alochym@123
        -   alochym/Alochym@123
    -   database name
        -   alochym
    -   https://linuxize.com/post/how-to-create-mysql-user-accounts-and-grant-privileges/
