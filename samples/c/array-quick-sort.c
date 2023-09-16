#include <stdio.h>

void swap(int *x, int *y)
{
    int temp = *y;
    *y = *x;
    *x = temp;
}

int partition(int *arr, int start, int end)
{
    int pivot = arr[end];

    int middle = start;
    int curr = start;
    while (curr < end)
    {
        if (arr[curr] <= pivot)
        {
            swap(&arr[curr], &arr[middle]);
            middle++;
        }
        
        curr++;
    }
    
    swap(&arr[middle], &arr[end]);

    return middle;
}

void quick_sort(int *arr, int start, int end)
{
    if (start < end)
    {
        // Choose last element to be the pivot, random index would be better?
        int pivot_index = partition(arr, start, end);

        // left
        quick_sort(arr, start, pivot_index - 1);
        quick_sort(arr, pivot_index + 1, end);
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
    quick_sort(arr, 0, 7);
    printf("\n");
    print_arr(arr, 8);

    return 0;
}
