package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

var exePath = "C:\\Users\\samsung\\Desktop\\vegas_starfruit\\Vegas.exe" //"C:\\Users\\samsung\\Desktop\\Starfruit_update\\Starfruit.exe"

func main() {

	fmt.Printf("🚀 프로그램 실행 중... [%s]\n", exePath)

	// 2. 프로그램 실행 (비동기 실행)
	cmd := exec.Command("powershell", "/c", "start", "-verb", "runas", exePath)

	err := cmd.Start()
	if err != nil {
		log.Fatalf("프로그램 실행 실패: %v", err)
	}

	// 3. 프로그램이 켜질 때까지 충분히 기다림 (컴퓨터 속도에 따라 조절)
	fmt.Println("⏳ 로딩 대기 중 (3초)...")
	time.Sleep(3 * time.Second)

	// 4. 아까 구한 좌표로 마우스 이동 및 클릭!
	targetX, targetY := 1492, 542 // 방금 구하신 좌표!

	fmt.Printf("🎯 좌표(%d, %d)로 이동하여 클릭합니다.\n", targetX, targetY)

	// 마우스 이동
	robotgo.Move(targetX, targetY)

	// 확실하게 하기 위해 잠깐 멈췄다가 클릭
	time.Sleep(500 * time.Millisecond)
	robotgo.Click("left") // 왼쪽 클릭

	fmt.Println("✅ 매크로 동작 완료!")

	select {}
}
