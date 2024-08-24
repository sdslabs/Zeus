#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

int main(int argc, char *argv[]) {
    int NUM_CHILDREN = argc > 1 ? atoi(argv[1]) : 5;
    pid_t pids[NUM_CHILDREN];
    int status;

    for (int i = 0; i < NUM_CHILDREN; ++i) {
        pid_t pid = fork();

        if (pid < 0) {
            // Fork failed
            perror("fork");
            exit(EXIT_FAILURE);
        } else if (pid == 0) {
            // Child process
            printf("Child %d (PID: %d) starting\n", i, getpid());
            
            // You can replace this with any command you want to run
            // For example, you could use execvp to run a shell command
            char *argv[] = {"/bin/sleep", "10", NULL}; // Example command: sleep for 10 seconds
            execvp(argv[0], argv);
            
            // execvp only returns on error
            perror("execvp");
            exit(EXIT_FAILURE);
        } else {
            // Parent process
            pids[i] = pid;
        }
    }

    // Parent process waits for all children to complete
    for (int i = 0; i < NUM_CHILDREN; ++i) {
        pid_t pid = waitpid(pids[i], &status, 0);
        if (pid < 0) {
            perror("waitpid");
        } else {
            if (WIFEXITED(status)) {
                printf("Child %d (PID: %d) exited with status %d\n", i, pid, WEXITSTATUS(status));
            } else {
                printf("Child %d (PID: %d) did not exit normally\n", i, pid);
            }
        }
    }

    printf("All child processes have completed.\n");
    return 0;
}
