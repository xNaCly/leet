import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

import helper.Helper;

public class m347 {
    public static void main(String[] args) {
        Integer[][] nums = {
                { 1, 1, 1, 2, 2, 3 },
                { 1 },
        };
        int[] k = { 2, 1 };
        Object[][] expected = {
                { 1, 2 },
                { 1 },
        };

        for (int i = 0; i < nums.length; i++) {
            int[] in = Arrays.stream(nums[i]).mapToInt(Integer::intValue).toArray();
            Helper.AssertArrayEquals(
                    Arrays.stream(topKFrequent(in, k[i])).mapToObj(e -> e).toArray(Integer[]::new),
                    expected[i]);
        }
    }

    public static int[] topKFrequent(int[] nums, int k) {
        HashMap<Integer, Integer> c = new HashMap<>();

        for (int i : nums) {
            if (c.containsKey(i)) {
                c.put(i, c.get(i).intValue() + 1);
            } else {
                c.put(i, 1);
            }
        }

        ArrayList<Integer> res = new ArrayList<>();
        c.entrySet().stream().sorted((a, b) -> b.getValue().compareTo(a.getValue())).forEach(e -> res.add(e.getKey()));
        return res.stream().mapToInt(i -> Integer.valueOf(i)).limit(k).toArray();
    }
}
