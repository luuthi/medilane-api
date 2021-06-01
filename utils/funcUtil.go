package utils

//func CheckPermission(context echo.Context, server *s.Server, handlerFunc echo.HandlerFunc, permRepo repositories.PermissionRepository) error  {
//	key := context.Request().Method + strings.Replace(context.Request().RequestURI, server.Permission.BaseURL, "", 1)
//	requiredScope := server.Permission.Route[key]
//	server.Logger.Infof("requiredScope: %v", requiredScope)
//
//	token, err := VerifyToken(context.Request(), server)
//	if err != nil {
//		return context.JSON(http.StatusUnauthorized, responses.Data{
//			Code:    http.StatusUnauthorized,
//			Message: "invalid token",
//		})
//	}
//
//	claims, ok := token.Claims.(*token2.JwtCustomClaims)
//	if !ok {
//		return context.JSON(http.StatusUnauthorized, responses.Data{
//			Code:    http.StatusUnauthorized,
//			Message: "invalid token",
//		})
//	}
//	userID := claims.ID
//	//permRepo := repositories.NewPermissionRepository(server.DB)
//	var rs []models2.Permission
//	permRepo.GetPermissionByUsername(&rs, userID)
//	var found bool
//	for _, perm := range rs {
//		if perm.PermissionName == requiredScope {
//			found = true
//		}
//	}
//	if !found {
//		return context.JSON(http.StatusForbidden, responses.Data{
//			Code:    http.StatusForbidden,
//			Message: "access denied",
//		})
//	}
//	return handlerFunc(context)
//}
//
//func ExtractToken(r *http.Request) string {
//	bearToken := r.Header.Get("Authorization")
//	//normally Authorization the_token_xxx
//	strArr := strings.Split(bearToken, " ")
//	if len(strArr) == 2 {
//		return strArr[1]
//	}
//	return ""
//}
//
//func VerifyToken(r *http.Request, server *s.Server) (*jwt.Token, error) {
//	tokenString := ExtractToken(r)
//	token, err := jwt.ParseWithClaims(tokenString, &token2.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		//Make sure that the token method conform to "SigningMethodHMAC"
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(server.Config.Auth.AccessSecret), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return token, nil
//}
