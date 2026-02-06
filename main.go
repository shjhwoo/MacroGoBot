package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

var exePath = "C:\\Users\\samsung\\Desktop\\Starfruit_update\\Starfruit.exe"

func main() {

	fmt.Println("🚀 프로그램 실행 중...")

	exec.Command("powershell", "/c", "start", "-verb", "runas", exePath).Start()

	fmt.Println("⏳ 프로그램 로딩 대기 중 (5초)...")
	time.Sleep(5 * time.Second)

	Login()

	select {}
}

func Login() {
	// 2. 포커스 잡기 (화면 중앙 아무데나 한번 클릭)
	// 창이 떴는데 포커스가 다른데 가 있을 수 있으므로 안전장치
	robotgo.MoveClick(960, 540)
	time.Sleep(1 * time.Second)

	// 3. [핵심] 탭 키로 이동 (직접 세어본 횟수만큼 반복)
	// 예: 아이디 창까지 탭 3번이 필요하다면
	fmt.Println("🎹 입력창 찾아가는 중...")

	robotgo.KeyTap("tab")
	time.Sleep(500 * time.Millisecond)

	robotgo.KeyTap("tab")
	time.Sleep(500 * time.Millisecond)

	robotgo.KeyTap("tab")
	time.Sleep(500 * time.Millisecond)

	// 4. 입력 및 로그인
	fmt.Println("✍️ 아이디/비번 입력")
	robotgo.Type("shjhwoo@trustnhope.com")

	robotgo.KeyTap("tab") // 비번창으로 이동
	time.Sleep(500 * time.Millisecond)

	robotgo.Type("Qwe123!@#") // 비번 입력
	time.Sleep(500 * time.Millisecond)

	robotgo.KeyTap("enter") // 엔터 쳐서 로그인!

	fmt.Println("✅ 로그인 시도 완료")
}
