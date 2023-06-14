import helper.Helper;

public class m238 {
    public static void main(String[] args) {
        int[][] in = new int[][] {
                { 1, 2, 3, 4 },
                { -1, 1, 0, -3, 3 },
        };
        int[][] out = new int[][] {
                { 24, 12, 8, 6 },
                { 0, 0, 9, 0, 0 },
        };
        for (int i = 0; i < in.length; i++) {
            int[] r = productExceptSelf(in[i]);
            Helper.AssertIntArrayEquals(r, out[i]);
        }
    }

    public static int[] productExceptSelf(int[] nums) {
        int[] arr = new int[nums.length];
        int right = 1, left = 1;

        for (int i = 0; i < nums.length; i++) {
            arr[i] = left;
            left *= nums[i];
        }

        for (int i = nums.length - 1; i >= 0; i--) {
            arr[i] *= right;
            right *= nums[i];
        }
        return arr;
    }
}
