package main

import "fmt"

type tv struct {
}

func (t *tv) on() {
	fmt.Println("打开电视")
}

func (t *tv) off() {
	fmt.Println("关闭电视")
}

type gc struct {
}

func (t *gc) on() {
	fmt.Println("打开游戏机")
}

func (t *gc) off() {
	fmt.Println("关闭游戏机")
}

type lamp struct {
}

func (t *lamp) on() {
	fmt.Println("开灯")
}

func (t *lamp) off() {
	fmt.Println("关灯")
}

type microphone struct {
}

func (t *microphone) on() {
	fmt.Println("打开麦克风")
}

func (t *microphone) off() {
	fmt.Println("关闭麦克风")
}

type audio struct {
}

func (t *audio) on() {
	fmt.Println("打开音响")
}

func (t *audio) off() {
	fmt.Println("关闭音响")
}

type projector struct {
}

func (t *projector) on() {
	fmt.Println("打开投影仪")
}

func (t *projector) off() {
	fmt.Println("关闭投影仪")
}

type facade struct {
	tv         *tv
	gc         *gc
	lamp       *lamp
	microphone *microphone
	audio      *audio
	projector  *projector
}

func (f facade) moviePattern() {
	f.tv.on()
	f.microphone.on()
	f.audio.on()
	f.lamp.off()
}
func (f facade) ktvPattern() {
	f.microphone.on()
	f.audio.on()
	f.lamp.off()
}

/**
优点：1、松耦合
缺点:如果子类发生改变那么就可能违反开闭原则
*/
func main() {
	fc := &facade{lamp: new(lamp), tv: new(tv), gc: new(gc), microphone: new(microphone), audio: new(audio), projector: new(projector)}
	fc.ktvPattern()
}
