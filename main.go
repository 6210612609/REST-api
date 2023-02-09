package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"example.com/controller"
	"example.com/service"
	"github.com/gin-gonic/gin"
)

var (
	assignmentService    service.AssignmentService       = service.New()
	assignmentController controller.AssignmentController = controller.New(assignmentService)
	problemService       service.ProblemService          = service.NewProblem()
	problemController    controller.ProblemController    = controller.NewProblem(problemService)
)


func run_file(path string, filename string, language string) {
	
	fmt.Println("expect => C:/Users/npt/project/storage/")
	fmt.Println("get => "+path)
	fmt.Println("expect => hello.java")
	fmt.Println("get => "+filename)
	fmt.Println("expect => java")
	fmt.Println("get => "+language)

	if language == "java"{
		cmd0, err0 := exec.Command("javac", path+filename+".java").CombinedOutput()
	if err0 != nil {
		fmt.Println(err0)
	}
	fmt.Println(string(cmd0))

	cmd1, err1 := exec.Command("java", "-cp", path, filename).CombinedOutput()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(string(cmd1))
	}else if language == "py" {
		cmd2, err2 := exec.Command("python3", path+filename+".py").CombinedOutput()
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(string(cmd2))
        
    }else{
		fmt.Println("file not found")
	}

}

func main() {
	server := gin.Default()

	server.GET("/compile", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, assignmentController.FindAll())

	})
	server.POST("/compile", func(ctx *gin.Context) {
		ctx.JSON(200, assignmentController.Save(ctx))

		file, _ := ctx.FormFile("source")

		// Upload the file to specific dst.
		storage_path := "C:/Users/npt/project/storage/"
		dst := storage_path + file.Filename
		ctx.SaveUploadedFile(file, dst)

		filename := strings.Split(file.Filename, ".")

		run_file(storage_path,filename[0],filename[1])
	})

	server.GET("/problem", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, problemController.FindAllProblem())
	})
	server.POST("/problem", func(ctx *gin.Context) {
		ctx.JSON(200, problemController.Commit(ctx))
	})

	server.Run(":8081")
}