package com.datatist1118;

import java.util.Scanner;

/*
 * ���ʽ��ȡֵ��byte,short,int,char
 * jdk5�Ժ������ö��
 * jdk7�Ժ������String
 * ��ݼ����Դ�����и�ʽ��
 * ctrl+shift+f
 */
public class SwitchTest {
	public static void main(String[] args) {

		Scanner sc = new Scanner(System.in);
		System.out.println("������һ��������1-7��");
		int weekDay = sc.nextInt();
		switch (weekDay) {
		/*
		 * 
		 * break;���"���Ǳ����"��
		 * �����д�����һ��case��Ӧ��ֵ�ɹ���
		 * ���ڲ�û��break��䣬
		 * ��ô����������(���ٽ���caseƥ��)
		 * �ļ�������ִ������case�е���䣬
		 * ֱ������break;�����ߵ���switch������
		 */
		case 1:
			System.out.println("����һ");
		case 2:
			System.out.println("���ڶ�");
		case 3:
			System.out.println("������");
		case 4:
			System.out.println("������");
			break;
		case 5:
			System.out.println("������");
			break;
		case 6:
			System.out.println("������");
			break;
		default:
			/*
			 * default�����е�ֵ���ͱ��ʽ��ƥ�䣬��ִ��default��Ӧ������
			 */
			System.out.println("������");
			break;
		}
	}
}
