package com.datatist1118;

import java.util.Scanner;

/*三元运算符练习之比较两个整数是否相同*/
public class Study2 {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.println("请输入一个整数：");
		int a = sc.nextInt();
		int b = 10;
		boolean flag = (a == b)?true:false;
		System.out.println(flag);
		sc.close();
	}
}
