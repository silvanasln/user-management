package main

import (
	"fmt"

	"github.com/silvanasln/user-management/repository"
	"github.com/silvanasln/user-management/service"
)
    
  


func main() {
    // رپوزیتوری و سرویس رو راه‌اندازی می‌کنیم
    repo := repository.NewInMemoryUserRepository()
    userService := service.NewUserService(repo)

    // ایجاد کاربر
    user, err := userService.CreateUser("name", "name@example.com")
    if err != nil {
        fmt.Println("خطا در ایجاد کاربر:", err)
        return
    }
    fmt.Printf("کاربر ایجاد شد: %+v\n", user)

    // دریافت همه کاربران
    users, err := userService.GetAllUsers()
    if err != nil {
        fmt.Println("خطا در دریافت کاربران:", err)
        return
    }
    fmt.Println("همه کاربران:", users)

    // به‌روزرسانی کاربر
    updatedUser, err := userService.UpdateUser(user.ID, "new name", "new.new@example.com")
    if err != nil {
        fmt.Println("خطا در به‌روزرسانی:", err)
        return
    }
    fmt.Printf("کاربر به‌روزرسانی شد: %+v\n", updatedUser)

    // دریافت کاربر با ID
    fetchedUser, err := userService.GetUserByID(user.ID)
    if err != nil {
        fmt.Println("خطا در دریافت کاربر:", err)
        return
    }
    fmt.Printf("کاربر دریافت‌شده: %+v\n", fetchedUser)

    // حذف کاربر
    err = userService.DeleteUser(user.ID)
    if err != nil {
        fmt.Println("خطا در حذف کاربر:", err)
        return
    }
    fmt.Println("کاربر حذف شد")
}

