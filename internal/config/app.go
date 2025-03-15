func BootStrap(config *BootstrapConfig) {
	//setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	postRepository := repository.NewPostRepository(config.Log)
	jobseekerProfileRepository := repository.NewJobseekerProfileRepository(config.Log)
	companyProfileRepository := repository.NewCompanyProfileRepository(config.Log)
	
	//setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	postUseCasse := usecase.NewPostUseCase(config.DB, config.Log, config.Validate, postRepository)
	profileUseCase := usecase.NewProfileUseCase(
		config.DB, 
		config.Log, 
		config.Validate, 
		jobseekerProfileRepository, 
		companyProfileRepository,
		userRepository,
	)
	
	//setup controllers
	userController := http.NewUserController(userUseCase, config.Log)
	postController := http.NewPostController(postUseCasse, config.Log)
  profileController := http.NewProfileController(profileUseCase, config.Log)
	//setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)
	
	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		PostController:    postController,
		ProfileController: profileController,
		AuthMiddleware:    authMiddleware,
	}
	routeConfig.Setup()
}