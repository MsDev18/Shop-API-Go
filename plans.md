validator ✅
service ✅
migrator ✅
migrations ✅
entity ✅
pkg & richerror ✅
repository ✅
docker compose ✅
config ✅
dto ✅
server ✅
router ✅
handler ✅
middleware ✅
response ✅

project bussinus logic convetions ✅
detect query in database for example run query 
isExist , err isExistUser := s.repository.ExistUserByPhoneNumber()
if !isExist {
    // more
}


api -
|____ health
|        |_____ GET /health-check
|
|____ auth
|        |_____ POST /auth/send-otp
|        |_____ POST /auth/check-otp
|        |_____ GET /auth/me
|        |_____ POST /auth/logout
|        |_____ POST /auth/refresh-token
|
|____ user
|        |_____ POST /user/profile
|        |_____ PATCH /user/update-profile
|        |_____ PATCH /user/change-password
|
|____ category