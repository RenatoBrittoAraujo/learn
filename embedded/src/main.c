#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <termios.h>
#include <stdlib.h>
#include <string.h>

#define IINTSOL 0xA1
#define IFLTSOL 0xA2
#define ICHRSOL 0xA3

#define IINTSND 0xB1
#define IFLTSND 0xB2
#define ICHRSND 0xB3

struct Message
{
    unsigned char *input_buffer;
    int input_buffer_size;
    unsigned char *output_buffer;
    int output_buffer_size;
};
typedef struct Message Message;
void transact_buffer_to_UART(Message *msgbuf);

Message get_solicition_message(int cod, char matricula[4])
{
    Message message;
    message.input_buffer = (unsigned char *)malloc(sizeof(unsigned char) * 5);
    message.input_buffer[0] = cod;
    for (int i = 0; i < 4; i++)
    {
        message.input_buffer[i + 1] = matricula[i];
    }
    message.input_buffer_size = 5;
    return message;
}

Message get_send_integer_message(char matricula[4], int value)
{
    Message message;
    message.input_buffer = (unsigned char *)malloc(sizeof(unsigned char) * 9);
    message.input_buffer[0] = IINTSND;
    message.input_buffer[1] = (value >> 24) & 0xFF;
    message.input_buffer[2] = (value >> 16) & 0xFF;
    message.input_buffer[3] = (value >> 8) & 0xFF;
    message.input_buffer[4] = value & 0xFF;
    for (int i = 0; i < 4; i++)
    {
        message.input_buffer[i + 5] = matricula[i];
    }
    message.input_buffer_size = 9;
    return message;
}

Message get_send_float_message(char matricula[4], float value)
{
    Message message;
    message.input_buffer = (unsigned char *)malloc(sizeof(unsigned char) * 9);
    message.input_buffer[0] = IFLTSND;
    union
    {
        float value;
        unsigned char bytes[4];
    } float2bytearray;
    float2bytearray.value = value;
    memcpy(message.input_buffer + 1, float2bytearray.bytes, 4);
    for (int i = 0; i < 4; i++)
    {
        message.input_buffer[i + 5] = matricula[i];
    }
    message.input_buffer_size = 9;
    return message;
}

Message get_send_string_message(char matricula[4], char *msg)
{
    int len = strlen(msg);
    if (len > 255)
    {
        printf("Mano sinceramente vc quer msm guardar mais de 255 caracteres? Tipo, complicado tlg...\n");
    }
    Message message;
    message.input_buffer = (unsigned char *)malloc(sizeof(unsigned char) * (len + 6));
    message.input_buffer[0] = ICHRSND;
    message.input_buffer[1] = len;
    for (int i = 0; i < len; i++)
    {
        message.input_buffer[i + 2] = msg[i];
    }
    for (int i = 0; i < 4; i++)
    {
        message.input_buffer[i + len + 2] = matricula[i];
    }
    message.input_buffer_size = len + 6;
    return message;
}

int main(int argc, char *argv[])
{
    if (argc < 2)
    {
        printf("Usage: %s <last 4 digits of your 'matricula'>\n", argv[0]);
        return 1;
    }

    char *strmsg = "goiás é saudade em tudo que falo, às vezes me calo por esta razão";

    Message get_int = get_solicition_message(IINTSOL, argv[1]);
    Message get_float = get_solicition_message(IFLTSOL, argv[1]);
    Message get_string = get_solicition_message(ICHRSOL, argv[1]);
    Message send_int = get_send_integer_message(argv[1], 69);
    Message send_float = get_send_float_message(argv[1], 42.0);
    Message send_string = get_send_string_message(argv[1], strmsg);

    Message msgs[6] = {
        get_int,
        get_float,
        get_string,
        send_int,
        send_float,
        send_string};

    for (int i = 0; i < 6; i++)
    {
        printf("===== UART TRANSACTION =====\n");
        transact_buffer_to_UART(&msgs[i]);
        printf("=== UART TRANSACTION END ===\n");
        printf("Sent Buffer: %s\n", msgs[i].input_buffer);
        printf("Got Buffer: %s\n", msgs[i].output_buffer);
    }

    return 0;
}

void transact_buffer_to_UART(Message *msgbuf)
{

    int uart0_filestream = -1;

    uart0_filestream = open("/dev/serial0", O_RDWR | O_NOCTTY | O_NDELAY); // Open in non blocking read/write mode
    if (uart0_filestream == -1)
    {
        printf("Erro - Não foi possível iniciar a UART.\n");
    }
    else
    {
        printf("UART inicializada!\n");
    }
    struct termios options;
    tcgetattr(uart0_filestream, &options);
    options.c_cflag = B9600 | CS8 | CLOCAL | CREAD; //<Set baud rate
    options.c_iflag = IGNPAR;
    options.c_oflag = 0;
    options.c_lflag = 0;
    tcflush(uart0_filestream, TCIFLUSH);
    tcsetattr(uart0_filestream, TCSANOW, &options);

    unsigned char *tx_buffer = (*msgbuf).input_buffer;
    unsigned char *p_tx_buffer = (*msgbuf).input_buffer + (*msgbuf).input_buffer_size;

    if (uart0_filestream != -1)
    {
        printf("Writing to UART ...");
        int count = write(uart0_filestream, &tx_buffer[0], (p_tx_buffer - &tx_buffer[0]));
        if (count < 0)
        {
            printf("UART TX error\n");
        }
        else
        {
            printf("Write completed.\n");
        }
    }

    sleep(1);

    //----- CHECK FOR ANY RX BYTES -----
    if (uart0_filestream != -1)
    {
        // Read up to 255 characters from the port if they are there
        unsigned char *rx_buffer = (unsigned char *)malloc(sizeof(unsigned char) * 256);
        int rx_length = read(uart0_filestream, (void *)rx_buffer, 255); // Filestream, buffer to store in, number of bytes to read (max)
        if (rx_length < 0)
        {
            printf("Erro na leitura.\n"); // An error occured (will occur if there are no bytes)
        }
        else if (rx_length == 0)
        {
            printf("Nenhum dado disponível.\n"); // No data waiting
        }
        else
        {
            // Bytes received
            rx_buffer[rx_length] = '\0';
            printf("%i Bytes lidos : %s\n", rx_length, rx_buffer);
        }
        (*msgbuf).output_buffer = rx_buffer;
        (*msgbuf).output_buffer_size = rx_length;
    }

    close(uart0_filestream);
}
