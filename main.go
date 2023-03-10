package main

import (
	"fmt"
	"io"
	"log"
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
	input_type := "keyboard"
	// inputs := strings.Split(input, ",")

	if input_type == "args" {
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
			input1 := "11"
			input2 := "12"
			input3 := "13"
			input4 := "14"
			cmd, err := exec.Command("python3", path+filename+".py", input1,input2,input3,input4).CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Print(string(cmd))
		} else {
			fmt.Println("file not found")
		}
	} else if input_type == "keyboard" {
		if language == "java" {
			cmd := exec.Command("python3", path+filename+".java")
			stdin, err := cmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}

			// input
			input := "12" + "\n" + "13" + "\n" + "14"

			go func() {
				defer stdin.Close()
				io.WriteString(stdin, input)
			}()

			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("----------------------------------------")
			fmt.Printf("%s", out)
			fmt.Println("----------------------------------------")
		} else if language == "python" {
			cmd := exec.Command("python3", path+filename+".py")
			stdin, err := cmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}

			// input
			input := "12" + "\n" + "13" + "\n" + "14"

			go func() {
				defer stdin.Close()
				io.WriteString(stdin, input)
			}()

			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("----------------------------------------")
			fmt.Printf("%s", out)
			fmt.Println("----------------------------------------")
		} else {
			fmt.Println("file not found")
		}

	} else {
		fmt.Println("wrong input")
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
