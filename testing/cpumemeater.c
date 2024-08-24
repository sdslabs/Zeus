#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Usage: %s <memory_size_in_MB>\n", argv[0]);
        return 1;
    }

    // Output the PID
    printf("PID: %d\n", getpid());

    long long int memory_size = atoll(argv[1]);
    long long int num_elements = memory_size * 1024 * 1024 / sizeof(int);

    int *memory = (int *)malloc(num_elements * sizeof(int));
    if (memory == NULL) {
        printf("Memory allocation failed!\n");
        return 1;
    }

    // Fill the allocated memory with some data
    for (long long int i = 0; i < num_elements; i++) {
        memory[i] = i;
    }

    return 0;
}