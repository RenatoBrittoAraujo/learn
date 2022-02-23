/* ******************************/
/* FGA/Eng. Software/ FRC       */
/* Prof. Fernando W. Cruz       */
/* Codigo: tcpClient2.c         */
/* ******************************/

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h> // close
#include <math.h>   // sqrt
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

#define MAX_SIZE 80

struct server_res
{
    float max;
    float min;
};
typedef struct server_res server_res;

struct thread_args
{
    int index;
    float *inp_arr;
    int arr_size;
    char *ip;
    char *port;
    server_res ***responses;
    int *responses_count;
    int expected_responses;
};
typedef struct thread_args thread_args;

float *gen_arr(int size)
{
    float *arr = (float *)malloc(sizeof(float) * size);
    int i;
    for (i = 0; i < size; i++)
    {
        arr[i] = (i - size / 2.0) * (i - size / 2.0);
        arr[i] = sqrt(arr[i]);
    }
    return arr;
}

server_res calculate_server_res(char *bufin)
{
    union
    {
        float f;
        char c[4];
    } min_o, max_o;
    for (int i = 0; i < 4; i++)
    {
        if (bufin[i] == '*')
            min_o.c[i] = 0;
        else
            min_o.c[i] = bufin[i];
    }
    for (int i = 0; i < 4; i++)
    {
        if (bufin[4 + i] == '*')
            max_o.c[i] = 0;
        else
            max_o.c[i] = bufin[4 + i];
    }
    server_res ans = {
        min : min_o.f,
        max : max_o.f,
    };
    return ans;
}

server_res get_single_response(float *inp_arr, int arr_size, char *ip, char *port, char *thread_id)
{
    struct sockaddr_in ladoServ; /* contem dados do servidor 	*/
    int sd;                      /* socket descriptor              */
    int n, k;                    /* num caracteres lidos do servidor */
    char bufout[MAX_SIZE];       /* buffer de dados enviados  */
    char bufin[9];               /* buffer de dados recebidos  */

    memset((char *)&ladoServ, 0, sizeof(ladoServ)); /* limpa estrutura */
    memset((char *)&bufout, 0, sizeof(bufout));     /* limpa buffer */

    ladoServ.sin_family = AF_INET; /* config. socket p. internet*/
    ladoServ.sin_addr.s_addr = inet_addr(ip);
    ladoServ.sin_port = htons(atoi(port));

    printf("[CLIENT] [thread %s] --- connect to server %s:%s --- \n", thread_id, ip, port);
    /* Cria socket */
    sd = socket(AF_INET, SOCK_STREAM, 0);
    if (sd < 0)
    {
        fprintf(stderr, "[CLIENT] Criacao do socket falhou!\n");
        exit(1);
    }

    /* Conecta socket ao servidor definido */
    if (connect(sd, (struct sockaddr *)&ladoServ, sizeof(ladoServ)) < 0)
    {
        fprintf(stderr, "[CLIENT] Tentativa de conexao falhou!\n");
        exit(1);
    }

    int arr_p = 0;
    int c_p = 0;
    int package_count = 0, float_count = 0;

    // enquanto existem floats a serem enviados
    while (arr_p < arr_size)
    {
        memset(&bufout, 0x0, sizeof(bufout));
        c_p = 1;
        // bota o maximo de floats possíveis nesse buffer
        while (arr_p < arr_size && c_p + 4 < MAX_SIZE - 1)
        {
            union
            {
                float f;
                char c[4];
            } u;
            u.f = inp_arr[arr_p++];
            bufout[c_p++] = u.c[0];
            bufout[c_p++] = u.c[1];
            bufout[c_p++] = u.c[2];
            bufout[c_p++] = u.c[3];
        }
        // substitui todos os '0' por '*' para evitar lidar problemas com strlen()
        for (int i = 1; i < c_p; i++)
        {
            if (bufout[i] == 0)
            {
                bufout[i] = '*';
            }
        }
        // adiciona o tamanho do pacote no primeiro byte
        bufout[0] = c_p;
        bufout[c_p] = '\0';

        send(sd, &bufout, strlen(bufout), MSG_DONTWAIT); /* enviando dados ...  */

        // Recebe ACK
        recv(sd, bufin, 1, 0);
        c_p = 0;

        float_count += bufout[0] / 4;
        package_count++;
    }
    printf("[SERVER] [thread %s] Send %d packages, %d floats\n", thread_id, package_count, float_count);

    bufout[0] = 'R';
    bufout[1] = 'E';
    bufout[2] = 'S';
    bufout[3] = '\0';
    send(sd, &bufout, strlen(bufout), 0); /* enviando pedido de resposta ...  */

    memset(&bufin, 0x0, sizeof(bufin));
    int ret = recv(sd, bufin, 9, 0);

    close(sd);
    printf("[CLIENT] [thread %s] --- disconnect from server %s %s --- \n", thread_id, ip, port);
    return calculate_server_res(bufin);
}

pthread_cond_t cond1 = PTHREAD_COND_INITIALIZER;
pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;

void *get_divide_and_conquer_response_thread(void *args_input)
{
    thread_args args = (*(struct thread_args *)args_input);
    server_res *ans = malloc(sizeof(server_res));

    char thread_id_as_string[10];
    sprintf(thread_id_as_string, "%d", args.index);

    (*ans) = get_single_response(args.inp_arr, args.arr_size, args.ip, args.port, thread_id_as_string);

    pthread_mutex_lock(&lock);
    (*args.responses)[args.index] = ans;
    (*args.responses_count)++;

    // if final response, send signal for main thread to continue
    if (*args.responses_count == args.expected_responses)
    {
        pthread_cond_signal(&cond1);
    }

    // Kill thread, release lock
    free(args_input);
    pthread_mutex_unlock(&lock);
}

server_res get_divide_and_conquer_response(float *inp_arr, int arr_size, char *ip, char **portlist, int portcount)
{
    // separa tarefa em secções por servidor
    int sec_size = arr_size / portcount;
    float **sections = (float **)malloc(sizeof(float *) * portcount);
    int section_sizes[portcount];
    int pos = 0;
    for (int i = 0; i < portcount; i++)
    {
        int size = sec_size;
        if (i == portcount - 1)
        {
            size = arr_size - pos;
        }
        sections[i] = (float *)malloc(sizeof(float) * size);
        for (int j = pos; j < pos + sec_size; j++)
        {
            sections[i][j - pos] = inp_arr[j];
        }
        pos += size;
        section_sizes[i] = size;
    }

    server_res **responses = (server_res **)malloc(sizeof(server_res *) * portcount);
    int responses_count = 0;

    printf("[CLIENT] [thread main] Creating %d threads\n", portcount);

    // cria thread para cada secção
    for (int i = 0; i < portcount; i++)
    {
        pthread_t tid;
        thread_args *args = (thread_args *)malloc(sizeof(thread_args));
        (*args).index = i;
        (*args).inp_arr = sections[i];
        (*args).arr_size = section_sizes[i];
        (*args).ip = ip;
        (*args).port = portlist[i];
        (*args).responses = &responses;
        (*args).responses_count = &responses_count;
        (*args).expected_responses = portcount;
        pthread_create(&tid, NULL, get_divide_and_conquer_response_thread, (void *)args);
    }

    printf("[CLIENT] [thread main] Awaiting all threads to finish...\n");
    pthread_cond_wait(&cond1, &lock);
    printf("[CLIENT] [thread main] All threads finished\n");

    // unifica respostas
    server_res ans;
    ans.min = (*responses[0]).min;
    ans.max = (*responses[0]).max;
    for (int i = 1; i < portcount; i++)
    {
        if (ans.min > (*responses[i]).min)
            ans.min = (*responses[i]).min;
        if (ans.max < (*responses[i]).max)
            ans.max = (*responses[i]).max;
    }
    return ans;
}

int main(int argc, char *argv[])
{
    int arr_size = 10000;

    if (argc < 3)
    {
        printf("[CLIENT] uso correto: %s <ip_do_servidor> <porta_do_servidor> <...opcional: mais_portas_de_servidor>\n", argv[0]);
        exit(1);
    }
    char *ip = argv[1], *port = argv[2];

    float *inp_arr = gen_arr(arr_size);

    server_res ans;
    if (argc > 3)
    {
        char **portlist = (char **)malloc(sizeof(char *) * (argc - 2));
        for (int i = 2; i < argc; i++)
        {
            portlist[i - 2] = argv[i];
        }
        ans = get_divide_and_conquer_response(inp_arr, arr_size, ip, portlist, argc - 2);
    }
    else
    {
        ans = get_single_response(inp_arr, arr_size, ip, port, "main");
    }

    printf("[CLIENT] FINAL RESULT\n");
    printf("[CLIENT] min: %f\n", ans.min);
    printf("[CLIENT] max: %f\n", ans.max);
}
