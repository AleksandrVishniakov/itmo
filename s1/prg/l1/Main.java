import java.util.Random;

public class Main {
    public static boolean inNums(int e) {
        int[] nums = new int[] { 6, 8, 9, 10, 11, 13, 15 };
        for (int k = 0; k < nums.length; k++) {
            if (e == nums[k]) {
                return true;
            }
        }

        return false;
    }

    public static double calculate(short[] e, double[] x, int i, int j) {
        if (e[i] == 14) {
            return Math.asin(Math.pow(Math.cos(x[j]), 2));
        } else if (inNums(e[i])) {
            return Math.pow((0.5 + Math.cbrt(x[j])) / (x[j] * x[j]), 6);
        } else {
            return Math.pow(3.0 / 
                (2 + (1 - Math.pow(Math.cbrt(x[j]), Math.tan(x[j]) * (x[j] / 2 + 4))) / 
                (Math.cbrt(Math.log(Math.abs(x[j]))))), 2
            );
        }
    }

    public static void main(String[] args) {
        // 1
        short[] e = new short[15];
        for (int i = 0; i < e.length; i++) {
            e[i] = (short) (i + 6);
        }

        // 2
        double[] x = new double[12];
        Random random = new Random();
        for (int i = 0; i < x.length; i++) {
            x[i] = random.nextDouble(-10.0, 6.0);
        }

        // 3
        double[][] e1 = new double[15][12];
        for (int i = 0; i < e1.length; i++) {
            for (int j = 0; j < e1[i].length; j++) {
                e1[i][j] = calculate(e, x, i, j);
            }
        }

        // 4
        for (int i = 0; i < e1.length; i++) {
            for (int j = 0; j < e1[i].length; j++) {
                System.out.printf("%.4f\t", e1[i][j]);
            }
            System.out.println();
        }
    }
}
