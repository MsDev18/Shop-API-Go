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
|____ auth
        |_____ /auth/send-otp
        |_____ /auth/check-otp
        |_____ /auth/me
        |_____ /auth/logout
        |_____ /auth/refresh-token