package com.datatist1119;

import java.util.Random;

public class Demo01 {
	// 需求：在控制台输出所有的”水仙花数”
/*	public static void main(String[] args) {
		for (int i = 100; i < 1000; i++) {
			int oneNum = i % 10;
			int tenNum = i / 10 % 10;
			int hunNum = i / 10 / 10;
			if ((Math.pow(oneNum, 3) + Math.pow(tenNum, 3) + Math.pow(hunNum, 3)) == i) {
				System.out.println(i);
			}
		}
	}*/
	//7行8列星星
/*	public static void main(String[] args) {
		for(int i = 1;i<=7;i++){
			for(int j = 1;j<=8;j++){
				System.out.print("*");
			}
			System.out.println();
		}
	}*/
	//九九乘法表
/*	public static void main(String[] args) {
		for(int i = 1;i<=9;i++){
			for(int j =1;j<=i;j++){
				System.out.print(j+"*"+i+"="+j*i+"\t");
			}
			System.out.println();
		}
	}*/
	//random 1-100
	public static void main(String[] args) {
		Random i = new Random();
		int count = 0;
		for(int x = 1;x<=20;x++){
			int number = i.nextInt(100)+1;
			System.out.print(number+"\t");
			count+=1;
			if(count%5 == 0){
				System.out.println();
			}
		}	
	}
}
