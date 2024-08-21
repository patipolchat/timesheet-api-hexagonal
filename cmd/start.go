/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/adapter/handlers"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/adapter/repository"
	"github.com/patipolchat/timesheet-api-hexagonal/internal/core/services"
	"github.com/patipolchat/timesheet-api-hexagonal/pkg/customEcho"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:4000)/timesheetHexagonal"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		repo := repository.NewTimesheetRepository(db)
		service := services.NewTimesheetService(repo)
		handler := handlers.NewTimesheetHandler(service)
		e := echo.New()
		e.POST("/timesheet", handler.HandleCreateRequest)
		e.Use(customEcho.ContextMiddleware)
		e.Use(middleware.RequestID())
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		go func() {
			if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
				e.Logger.Fatal("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
