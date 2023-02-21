package main

import (
	"fmt"
	// "io"
	// "log"
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

func run_file(path string, filename string, language string, input string) {

	// inputs := strings.Split(input, ",")

	if language == "java" {
		cmd0, err0 := exec.Command("javac", path+filename+".java").CombinedOutput()
		if err0 != nil {
			fmt.Println(err0)
		}
		fmt.Println(string(cmd0))

		cmd1, err1 := exec.Command("java", "-cp", path, filename, input).CombinedOutput()
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println(string(cmd1))

	} else if language == "python" {

		//args
		prg := "python3"

		arg1 := "C:/Users/npt/project/REST-api/test2.py"
		arg2 := "2"
		arg3 := "3"

		cmd := exec.Command(prg, arg1, arg2, arg3)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print(string(stdout))



		//keyboard input

		// cmd := exec.Command("python3", "C:/Users/npt/project/REST-api/test2.py")
		// stdin, err := cmd.StdinPipe()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// // input
		// input := "hello "+"\n"+"world "+"\n"+"end "

		// go func() {
		// 	defer stdin.Close()
		// 	io.WriteString(stdin, input)
		// }()

		// out, err := cmd.CombinedOutput()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println("----------------------------------------")
		// fmt.Printf("%s", out)
		// fmt.Println("----------------------------------------")

	} else {
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
		full_filename := strings.Split(file.Filename, ".")
		filename := full_filename[0]
		language := ctx.Request.FormValue("language")
		input := ctx.Request.FormValue("input")

		// Upload the file to specific dst.
		storage_path := "C:/Users/npt/project/storage/"
		dst := storage_path + file.Filename
		ctx.SaveUploadedFile(file, dst)

		run_file(storage_path, filename, language, input)
	})

	server.GET("/problem", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, problemController.FindAllProblem())
	})
	server.POST("/problem", func(ctx *gin.Context) {
		ctx.JSON(200, problemController.Commit(ctx))
	})

	server.Run(":8081")
}
