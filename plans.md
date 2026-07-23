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

solution 1 
error handling with richerror for example 
destruct richerror for check bussinus login for err not found 

solution 2
detect query in database for example run query 
isExist , err isExistUser := s.repository.ExistUserByPhoneNumber()
if !isExist {
    // more
}
