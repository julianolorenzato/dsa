#include <stdio.h>
#include <stdbool.h>

void swap(int *x, int *y)
{
    int temp = *y;
    *y = *x;
    *x = temp;
}

int *bubble_sort(int *arr, int arrSize)
{
    for (int i = 0; i < arrSize; i++)
    {
        for (int j = 0; j < arrSize - i - 1; j++)
        {
            if (arr[j] > arr[j + 1])
            {
                swap(&arr[j], &arr[j + 1]);
            }
        }
    }
}

// This version of bubble sort contains a boolean flag, that indicates
// if were not swaps, if this happen, the arr is already sorted, and it breaks the outter loop
int *bubble_sort_with_flag(int *arr, int arrSize)
{
    bool is_sorted = true;

    for (int i = 0; i < arrSize; i++)
    {
        for (int j = 0; j < arrSize - i - 1; j++)
        {
            if (arr[j] > arr[j + 1])
            {
                is_sorted = false;
                swap(&arr[j], &arr[j + 1]);
            }
        }

        if (is_sorted)
            break;
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
    bubble_sort(arr, 8);
    printf("\n");
    print_arr(arr, 8);

    return 0;
}
