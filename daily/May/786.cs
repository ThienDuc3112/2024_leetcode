namespace KthSmallestPrimeFraction;

public class Solution
{
    public int[] KthSmallestPrimeFraction(int[] arr, int k)
        // Editorial solution
    {
        double left = 0, right = 1.0;
        int n = arr.Length;

        while (left < right)
        {
            double mid = (left + right) / 2;

            double maxFraction = 0;
            int countSmallerFraction = 0, numeratorIdx = 0, denominatorIdx = 0;
            int j = 1;

            for (int i = 0; i < arr.Length - 1; i++)
            {
                while (j < n && arr[i] >= mid * arr[j]) j++;
                /* if (j < n) Console.WriteLine($"mid: {mid}\tarr[i]/arr[j]: {arr[i]/arr[j]}"); */
                countSmallerFraction += n - j;
                if (j == n) break;
                double fraction = (double) arr[i] / arr[j];

                if (fraction > maxFraction)
                    (maxFraction, numeratorIdx, denominatorIdx) = (fraction, i, j);
                /* Console.WriteLine($"maxFraction", numeratorIdx, denominatorIdx); */
            }

            if (countSmallerFraction == k) return new int[2] { arr[numeratorIdx], arr[denominatorIdx] };
            else if (countSmallerFraction > k) right = mid;
            else left = mid;
        }
        return new int[2];
    }

    // Self make (merge m array of sorted array)
    /* { */
    /*     PriorityQueue<int[], double> pq = new(); */
    /*     int[][] matrix = new int[arr.Length * arr.Length / 2][]; */
    /*     int idx = 0; */
    /*     int[] pointers = new int[arr.Length]; */
    /*     for (int i = 0; i < arr.Length; i++) */
    /*     { */
    /*         pointers[i] = arr.Length - 1; */
    /*         if (pointers[i] > i)  */
    /*             pq.Enqueue(new int[3] { arr[i], arr[pointers[i]], i }, ((double)arr[i]) / arr[pointers[i]--]); */
    /*     } */
    /*     while(pq.Count > 0) { */
    /*         var cur = pq.Dequeue(); */
    /*         if(pointers[cur[2]] > cur[2]){ */
    /*             pq.Enqueue(new int[3] { arr[cur[2]], arr[pointers[cur[2]]], cur[2] }, ((double)arr[cur[2]]) / arr[pointers[cur[2]]--]); */
    /*         } */
    /*         matrix[idx++] = new int[2]{cur[0], cur[1]}; */
    /*     } */
    /*     return matrix[k-1]; */
    /* } */
}
