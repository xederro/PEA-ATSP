import csv
import os
import sys
import matplotlib.pyplot as plt
import numpy as np
from matplotlib.ticker import FuncFormatter
from scipy.interpolate import make_interp_spline


def read_csv(file_path):
    data = {}
    with open(file_path, "r") as file:
        reader = csv.reader(file, delimiter=";")
        for row in reader:
            algorithm, size, good, todo, time_in_nanoseconds = row
            size = int(size)
            time_in_nanoseconds = int(time_in_nanoseconds)
            if algorithm not in data:
                data[algorithm] = {
                    "size": [],
                    "time_in_nanoseconds": [],
                }
            data[algorithm]["size"].append(size)
            data[algorithm]["time_in_nanoseconds"].append(time_in_nanoseconds)
    return data


def plot_specific_data(data, save_folder):
    # Define the combinations of data entries to plot
    specific_plots = [
        {
            "algorithm": None,
        },
        {
            "algorithm": "Brute Force",
        },
        {
            "algorithm": "Branch and Bound",
        },
        {
            "algorithm": "Dynamic programming",
        },
    ]

    for plot_spec in specific_plots:
        plt.figure(figsize=(10, 6))
        for algorithm, values in data.items():
            if plot_spec["algorithm"] is None or plot_spec["algorithm"] == algorithm:
                x = values["size"]
                y = values["time_in_nanoseconds"]
                X_Y_Spline = make_interp_spline(x, y)
                X_ = np.linspace(min(x), max(x), 500)
                Y_ = X_Y_Spline(X_)
                label = f"{algorithm} "
                plt.plot(X_, Y_, "-", label=label)
                plt.scatter(x, y)

        label = ""
        if plot_spec["algorithm"] is not None:
            label += f"{plot_spec['algorithm']} "
        if plot_spec["algorithm"] is None:
            label = "hewwo"

        label = label.strip()
        #         plt.title(label)

        plt.xlabel("Size")
        plt.ylabel("Time")
        plt.legend()
        plt.gca().yaxis.set_major_formatter(FuncFormatter(format_time_ticks))
        save_path = os.path.join(save_folder, f"{label}.svg")
        plt.savefig(save_path)
        plt.close()


def format_time_ticks(x, pos):
    units = ["ns", "μs", "ms", "s", "min", "hr", "day"]
    conversions = [1, 1000, 1000, 1000, 60, 60, 24]

    unit_index = 0
    while x >= 1000 and unit_index < len(units) - 1:
        x /= 1000
        unit_index += 1

    return f"{x:.2f} {units[unit_index]}"


def plot_data(data, save_folder):
    for name, values in data.items():
        plt.figure(figsize=(10, 6))
        x = values["size"]
        y = values["time_in_nanoseconds"]

        X_Y_Spline = make_interp_spline(x, y)
        X_ = np.linspace(min(x), max(x), 500)
        Y_ = X_Y_Spline(X_)
        plt.plot(X_, Y_, "-", color="blue")
        plt.scatter(x, y, color="blue")
        #         plt.title(name)
        plt.xlabel("Size")
        plt.gca().yaxis.set_major_formatter(FuncFormatter(format_time_ticks))
        save_path = os.path.join(save_folder, f"{name}.svg")
        plt.savefig(save_path)
        plt.close()


def main(file_path, save_folder):
    os.makedirs(save_folder, exist_ok=True)
    data = read_csv(file_path)
    # plot_data(data, save_folder)
    plot_specific_data(data, save_folder)


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python script.py <source_csv_file> <target_directory>")
        sys.exit(1)
    main(sys.argv[1], sys.argv[2])
