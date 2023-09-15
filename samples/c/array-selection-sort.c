#include <stdio.h>

void swap(int *x, int *y)
{
    int temp = *y;
    *y = *x;
    *x = temp;
}

int *selection_sort(int *arr, int arrSize)
{
    for (int i = 0; i < arrSize; i++)
    {
        for (int j = i + 1; j < arrSize; j++)
        {
            if (arr[i] > arr[j])
            {
                swap(&arr[i], &arr[j]);
            }
        }

        // another way, is better?
        // int i_min = i;

        // for (int j = i + 1; j < arrSize; j++)
        // {
        //     if (arr[j] < arr[i_min])
        //     {
        //         i_min = j;
        //     }
        // }
        
        // swap(&arr[i], &arr[i_min]);
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
    selection_sort(arr, 8);
    printf("\n");
    print_arr(arr, 8);

    return 0;
}
