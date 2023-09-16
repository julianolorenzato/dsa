void merge(int *arr, int start, int middle, int end)
{

    int left_length = middle - start + 1;
    int aux_left[left_length];

    int right_length = end - middle;
    int aux_right[right_length];

    for (int i = 0; i < left_length; i++)
    {
        aux_left[i] = arr[start + i];
    }

    for (int i = 0; i < right_length; i++)
    {
        aux_right[i] = arr[middle + 1 + i];
    }

    int i = 0, j = 0;
    for (int k = start; k <= end; k++)
    {
        if ((i < left_length) && (j >= right_length) || aux_left[i] <= aux_right[j])
        {
            arr[k] = aux_left[i];
            i++;
        }
        else
        {
            arr[k] = aux_right[j];
            j++;
        }
    }
}

void merge_sort(int *arr, int start, int end)
{
    if (end - start > 1)
    {
        int middle = (end + start) / 2;
        merge_sort(arr, start, middle);
        merge_sort(arr, middle + 1, end);

        merge(arr, start, middle, end);
    }
}