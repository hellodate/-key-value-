package main

 import(
	"gocode/testproject02/memCache/cache"
	"time"
 ) 

 func main(){
	cache := cache.NewMemCache()
	cache.SetMaxMemory("200MB")

	
	cache.Set("int" , 1  ,time.Second)
	cache.Set("bool" ,false, time.Second)
	cache.Set("date" ,map[string]interface{}{"a": 1}, time.Second)


	cache.Get("int")
	cache.Del("int")
	cache.Flush()
	cache.Keys()
 }