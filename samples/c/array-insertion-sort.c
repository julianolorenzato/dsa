#include <stdio.h>

void swap(int *x, int *y)
{
    int temp = *y;
    *y = *x;
    *x = temp;
}

int *insertion_sort(int *arr, int arrSize)
{
    for (int i = 0; i < arrSize; i++)
    {
        for (int j = i; j > 0; j--)
        {
            if (arr[j] < arr[j - 1])
            {
                swap(&arr[j], &arr[j - 1]);
            }
            else
            {
                break;
            }
        }

        // the same loop as above, but more concise
        // int j = i;
        // while (j > 0 && arr[j] < arr[j - 1])
        // {
        //     swap(&arr[j], &arr[j - 1]);
        //     j--;
        // }
    }
}

void print_arr(int *arr, int arrSize)
{
    for (int i = 0; i < arrSize; i++)
    {
        printf("[%d] - value: %d\n", i, arr[i]);
    }
}

int main(int argc, char const *argv[])
{
    int arr[8] = {2, 5, 4, 3, 6, 8, 1, 7};

    print_arr(arr, 8);
    insertion_sort(arr, 8);
    printf("\n");
    print_arr(arr, 8);

    return 0;
}
