package com.datatist1118;

import java.util.Scanner;

/*��Ԫ�������ϰ֮�Ƚ����������Ƿ���ͬ*/
public class Study2 {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		System.out.println("������һ��������");
		int a = sc.nextInt();
		int b = 10;
		boolean flag = (a == b)?true:false;
		System.out.println(flag);
		sc.close();
	}
}
