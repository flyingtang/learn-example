package function

func a() (result int){
	defer func(){
		result ++
	}()
	return 0
}

// 1

func b() (result int){
	t := 5
	defer func(){
		t = t + 5
	}()
	return t
}
// 5

func c() (result int){

	defer func(result int){

		result ++

	}(result)

	return 1
}
// 1