package app


func (receiver *server) InitRoutes() {
	mux := receiver.router.(*exactMux)
	mux.GET("/", receiver.handleBurgersList())

	mux.POST("/burgers/save", receiver.handleBurgersSave())
	mux.POST("/", receiver.handleBurgersList())

	mux.POST("/burgers/remove", receiver.handleBurgersRemove())

	mux.GET("/favicon.ico", receiver.handleFavicon())
}
