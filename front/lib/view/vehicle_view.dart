import 'package:flutter/material.dart';
import 'package:front/model/vehicle.dart';
import 'package:front/service/vehicle.dart';
import 'package:front/view/update_view.dart';

class VehiclesView extends StatefulWidget {
  const VehiclesView({super.key});

  @override
  State<VehiclesView> createState() => VehiclesViewState();
}

class _VehiclesViewState extends State<VehiclesView> {
  final VehicleService _vehicleService = VehicleService();
  late Future<List<VehicleModel>> _vehiclesFuture;

  @override
  void initState() {
    super.initState();
    _vehiclesFuture = _vehicleService.fetchData();
  }

  void _refreshVehicles() {
    setState(() {
      _vehiclesFuture = _vehicleService.fetchData();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: FutureBuilder<List<VehicleModel>>(
        future: _vehiclesFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else if (snapshot.hasError) {
            return Center(
              child: Text('Error: ${snapshot.error}'),
            );
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(
              child: Text('Nenhum usuário encontrado'),
            );
          } else {
            List<VehicleModel> vehicles = snapshot.data!;
            return ListView.builder(
              itemCount: vehicles.length,
              itemBuilder: (context, index) {
                return SizedBox(
                  height: 50,
                  child: Card(
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Name: ${vehicles[index].name}",
                        ),
                        Row(
                          children: [
                            IconButton(
                              onPressed: () async {
                                int id = vehicles[index].id;
                                await _vehicleService.deleteData(id: id);
                                _refreshVehicles();
                              },
                              icon: const Icon(Icons.delete),
                            ),
                            IconButton(
                              onPressed: () async {
                                int? id = vehicles[index].id;
                                await Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => UpdateView(
                                      vehicleId: id,
                                      name: vehicles[index].name,
                                    ),
                                  ),
                                );
                                _refreshVehicles();
                              },
                              icon: const Icon(Icons.edit),
                            ),
                          ],
                        ),
                      ],
                    ),
                  ),
                );
              },
            );
          }
        },
      ),
    );
  }
}
